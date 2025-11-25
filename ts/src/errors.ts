export class ApiError extends Error {
  readonly desc: string;

  constructor(desc: string, err: Error) {
    super(err.message);
    this.desc = desc;
    this.name = 'ApiError';
  }

  toJSON() {
    return {
      description: this.desc,
      error: this.message
    };
  }
}
