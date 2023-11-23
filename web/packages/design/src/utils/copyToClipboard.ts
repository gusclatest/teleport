/**
 * Copyright 2020 Gravitational, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * Copies text to clipboard.
 * Works only in browser, not in Electron.
 *
 * @param textToCopy the text to copy to clipboard
 */
export default async function copyToClipboard(
  textToCopy: string
): Promise<void> {
  try {
    await navigator.clipboard.writeText(textToCopy);
  } catch (err) {
    // Error can happen if the user denies clipboard permissions.
    window.prompt('Cannot copy to clipboard. Use ctrl/cmd + c', err);
  }
}
