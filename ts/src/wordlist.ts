import { promises as fs } from 'fs';
import * as path from 'path';

export interface WordList {
  addWord(word: string): Promise<void>;
  getWords(): Promise<string[]>;
}

export class FileWordList implements WordList {
  private readonly filename: string;

  constructor(filename: string) {
    this.filename = filename;
  }

  async addWord(word: string): Promise<void> {
    // mimic Go AddWord: append the raw word bytes (no newline handling here)
    await fs.appendFile(this.filename, word, { encoding: 'utf8' });
  }

  async getWords(): Promise<string[]> {
    // mimic Go GetWords conceptually: read all lines and trim whitespace
    const abs = path.resolve(this.filename);
    const data = await fs.readFile(abs, { encoding: 'utf8' });
    const lines: string[] = data.split('\n');
    return lines
      .map((s: string) => s.trim())
      .filter((s: string) => s.length > 0);
  }
}
