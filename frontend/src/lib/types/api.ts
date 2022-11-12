export interface APIResponseProps<T = Record<string, any>> {
	error: boolean;
	data: T;
	message?: string;
}
