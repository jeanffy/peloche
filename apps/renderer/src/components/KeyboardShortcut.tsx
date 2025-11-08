import { ReactElement, useEffect, useRef } from 'react';

export interface KeyboardShortcutProps {
  keyCodes: string[]; // https://www.quirksmode.org/js/keys.html
  shift?: boolean;
  control?: boolean;
  onTriggered: () => Promise<unknown> | unknown;
}

interface KeyboardShortcutInfo {
  id: string;
  keyCodes: string[];
  shift: boolean;
  control: boolean;
  callback: () => Promise<unknown> | unknown;
}

export default function KeyboardShortcut(props: KeyboardShortcutProps): ReactElement {
  const button = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    const shortcutId = getKeyboardShortcutId(props);

    activeKeyboardShortcuts.push({
      id: shortcutId,
      keyCodes: props.keyCodes,
      shift: props.shift ?? false,
      control: props.control ?? false,
      callback: () => {
        button.current?.click();
      },
    });

    return (): void => {
      const shortcutToRemove = activeKeyboardShortcuts.findIndex(s => s.id === shortcutId);
      if (shortcutToRemove !== -1) {
        activeKeyboardShortcuts.splice(shortcutToRemove, 1);
      }
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  async function handleButtonClick(): Promise<void> {
    const res = props.onTriggered();
    if (res instanceof Promise) {
      await res;
    }
  }

  return (
    <button
      aria-hidden="true"
      aria-description={`Keyboard shortcut for ${props.keyCodes.join(',')}`}
      ref={button}
      onClick={handleButtonClick}
      style={{ display: 'none' }}></button>
  );
}

const activeKeyboardShortcuts: KeyboardShortcutInfo[] = [];

document.addEventListener('keydown', async e => {
  if (e.repeat) return;

  for (const shortcut of activeKeyboardShortcuts) {
    if (shortcut.keyCodes.includes(e.code) && e.shiftKey === shortcut.shift && e.ctrlKey === shortcut.control) {
      const callbackReturn = shortcut.callback();
      if (callbackReturn instanceof Promise) {
        await callbackReturn;
      }
      e.preventDefault();
      break;
    }
  }
});

function getKeyboardShortcutId(shortcut: KeyboardShortcutProps): string {
  const shiftId = shortcut.shift === true ? '1' : '0';
  const controlId = shortcut.control === true ? '1' : '0';
  return `${shortcut.keyCodes.join()}${shiftId}${controlId}`;
}
