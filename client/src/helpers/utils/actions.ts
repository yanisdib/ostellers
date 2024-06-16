import { AnyAction } from '@reduxjs/toolkit';

import { FailedAction } from '../../interfaces/FailedAction';

export const isFailedAction = (action: AnyAction): action is FailedAction => {
	return action.type.endsWith('Failure');
};
