import { useEffect } from 'react';
import { useDispatch } from 'react-redux';

import { StoreState } from '@/interfaces/StoreState';
import { fetchArtbooksRequest } from '@/services/redux/slices/artbooks';

import { useAppSelector } from './useAppSelector';

export const useFetchArtbooks = () => {
	const dispatch = useDispatch();

	useEffect(() => {
		dispatch(fetchArtbooksRequest());
	});

	return useAppSelector((state) => state);
};
