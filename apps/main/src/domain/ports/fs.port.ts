export const FS_PORT_TOKEN = Symbol.for('FsPort');

export enum FsPortDirEntryType {
  File,
  Directory,
}

export interface FsPortDirEntry {
  type: FsPortDirEntryType;
  fullPath: string;
  parentDirPath: string;
  name: string;
  extension: string;
  extensionLowerCase: string;
  nameWithoutExtension: string;
}

export enum FsPortEncoding {
  Utf8,
  Latin1,
}

export interface FsPortReadFileOptions {
  encoding?: FsPortEncoding;
}

export interface FsPortWriteFileOptions {
  append?: boolean;
}

export interface FsPort {
  fileExists(filePath: string): Promise<boolean>;
  readDirectory(directoryPath: string): Promise<FsPortDirEntry[]>;
  readFile(filePath: string, options?: FsPortReadFileOptions): Promise<string>;
  renameFile(currentFilePath: string, newFilePath: string): Promise<void>;
  writeFile(filePath: string, data: string, options?: FsPortWriteFileOptions): Promise<void>;
  deleteFile(filePath: string): Promise<void>;
  createDirectory(directoryPath: string): Promise<void>;

  // transforms a file path (from the file system) to a URL path usable in the browser
  // handles app loaded with file:// and http:// protocols
  // when app is loaded with http protocol, the URL must be relative to the "public" folder of the renderer
  // if the file path is not, then a symbolic link must be created with a special name so that the URL can be relative
  // examples if public folder is '/app/renderer/public':
  //   - /app/renderer/public/foo/bar.txt -> foo/bar.txt
  //   - /tmp/foo/bar.txt -> <marker>/foo/bar.txt
  // TODO: more doc on this matter
  filePathToRendererUrl(filePath: string, replacements: Map<string, string>): Promise<string>;
}
