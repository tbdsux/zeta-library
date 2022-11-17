import { dev } from '$app/environment';

const apiUrl = dev ? 'http://localhost:8080' : '/api';

export { apiUrl };
