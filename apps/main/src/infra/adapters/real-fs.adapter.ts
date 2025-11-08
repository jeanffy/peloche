import fs from 'node:fs/promises';
import path from 'node:path';

import { FsPort, FsPortDirEntry, FsPortDirEntryType, FsPortWriteFileOptions } from '../../domain/ports/fs.port.js';
import { InjectedDependencies } from '../../types.js';
import { AppProtocol, INFRA_CONTEXT_PORT_TOKEN, InfraContextPort } from '../ports/infra-context.port.js';

export class RealFsAdapter implements FsPort {
  private contextPort: InfraContextPort;

  public constructor(dependencies: InjectedDependencies) {
    this.contextPort = dependencies[INFRA_CONTEXT_PORT_TOKEN];
  }

  public async fileExists(filePath: string): Promise<boolean> {
    let fileExists = false;
    try {
      await fs.stat(filePath);
      fileExists = true;
    } catch (error) {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      if ((error as any).code === 'ENOENT') {
        fileExists = false;
      } else {
        throw error;
      }
    }
    return fileExists;
  }

  public async readDirectory(directoryPath: string): Promise<FsPortDirEntry[]> {
    const entries = await fs.readdir(directoryPath, { withFileTypes: true });
    return entries
      .filter(e => e.isFile() || e.isDirectory())
      .map(e => ({
        type: e.isFile() ? FsPortDirEntryType.File : FsPortDirEntryType.Directory,
        fullPath: path.join(e.parentPath, e.name),
        parentDirPath: e.parentPath,
        name: e.name,
        extension: path.parse(e.name).ext,
        extensionLowerCase: path.parse(e.name).ext.toLowerCase(),
        nameWithoutExtension: path.parse(e.name).name,
      }));
  }

  public readFile(filePath: string): Promise<string> {
    return fs.readFile(filePath, { encoding: 'utf-8' });
  }

  public writeFile(filePath: string, data: string, options?: FsPortWriteFileOptions): Promise<void> {
    if (options?.append === true) {
      return fs.appendFile(filePath, data, { encoding: 'utf-8' });
    }
    return fs.writeFile(filePath, data, { encoding: 'utf-8' });
  }

  public async renameFile(currentFilePath: string, newFilePath: string): Promise<void> {
    return fs.rename(currentFilePath, newFilePath);
  }

  public async deleteFile(filePath: string): Promise<void> {
    return fs.unlink(filePath);
  }

  public async createDirectory(directoryPath: string): Promise<void> {
    await fs.mkdir(directoryPath, { recursive: true });
  }

  public async filePathToRendererUrl(filePath: string, replacements: Map<string, string>): Promise<string> {
    if (this.contextPort.protocol === AppProtocol.File) {
      return `file://${filePath}`;
    }

    // this case should happen only when running in dev mode or debugging, not in production
    // in that case, renderer (browser) will load files with http protocol, so file paths
    // need to be relative to the public folder
    // if filePath is relative to the public dir, just return the relative path as-is
    // if filePath is not relative, use the replacements defined in params

    const resolvedRendererPublicDir = path.resolve('./apps/renderer/public');
    const resolvedItemPath = path.resolve(filePath);

    if (resolvedItemPath.startsWith(resolvedRendererPublicDir)) {
      // here we transform a path like '/dev/apps/renderer/public/foo/bar.png' to 'foo/bar/png'
      return path.relative(resolvedRendererPublicDir, resolvedItemPath);
    }

    // here we transform a path like '/tmp/foo/bar.png' to '<marker>/foo/bar.png'
    // replacement is '/tmp' -> '<marker>'
    // '<marker>' symbolic link must exist in the apps/renderer/public directory
    let url = filePath;
    for (const [key, value] of replacements) {
      let dir = key;
      if (dir.endsWith(path.sep)) {
        dir = dir.slice(0, dir.length - 1);
      }
      url = url.replace(dir, value);
    }

    return url;
  }
}
