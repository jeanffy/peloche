import toast from 'react-hot-toast';

export function apiQuery<T>(fn: () => Promise<T>): () => Promise<T> {
  return () => {
    return fn().catch(error => {
      let message = 'Error';
      if (error instanceof Error) {
        message = error.message;
      }
      toast.error(message);
      throw error;
    });
  };
}
