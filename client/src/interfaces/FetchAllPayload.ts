export interface FetchAllPayload<T> {
	count: number;
	results: T;
	next?: string;
	previous?: string;
}
