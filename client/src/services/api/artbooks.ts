import axios, { AxiosResponse } from 'axios';

import { Artbook } from '@/interfaces/Artbook';
import { FetchAllPayload } from '@/interfaces/FetchAllPayload';

const ARTBOOKS_API_URI = `http://localhost:6060/v1/artbooks`;

export const fetchAllArtbooks = async (): Promise<
	FetchAllPayload<Artbook[]>
> => {
	try {
		const response: AxiosResponse<FetchAllPayload<Artbook[]>> =
			await axios.get(ARTBOOKS_API_URI);

		return response.data;
	} catch (error) {
		throw error;
	}
};
