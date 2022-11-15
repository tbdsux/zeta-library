import { PUBLIC_API_URL, PUBLIC_IS_DEV, PUBLIC_THEMOVIEDB_API_KEY } from '$env/static/public';

const isDev = PUBLIC_IS_DEV == 'true';
const apiUrl = isDev ? PUBLIC_API_URL : '/api';
const moviedbKey = PUBLIC_THEMOVIEDB_API_KEY ?? '';

export { apiUrl, isDev, moviedbKey };
