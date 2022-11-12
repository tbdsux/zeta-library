import { PUBLIC_API_URL, PUBLIC_IS_DEV } from '$env/static/public';

const isDev = PUBLIC_IS_DEV == 'true';
const apiUrl = isDev ? PUBLIC_API_URL : '/api';

export { apiUrl, isDev };
