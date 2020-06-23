import { all, fork } from 'redux-saga/effects'

import offersSagas from '../services/offers/sagas'

export default function* RootSaga(){
    yield all([
        fork(offersSagas)
    ])
}