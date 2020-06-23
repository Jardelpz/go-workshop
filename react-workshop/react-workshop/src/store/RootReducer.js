import { combineReducers } from 'redux'

import offersReducer from '../services/offers/reducers'

export default combineReducers({
    offers: offersReducer
})