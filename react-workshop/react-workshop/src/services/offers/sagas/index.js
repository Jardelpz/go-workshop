import { call, takeEvery } from 'redux-saga/effects'

import { OFFERS_REQUEST } from '../types'

import { getDebt } from '../../../api/offersApi'

function consoleLog(value) {
    console.log('Console:::', value)
}

function* loadOffers(){
    const data = yield call(getDebt, {})
    yield call(consoleLog, data)
}

export default function*(){
    yield takeEvery(OFFERS_REQUEST, loadOffers)
}