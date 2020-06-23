import React, { useEffect } from 'react'
import { connect } from 'react-redux'
import { loadOffers } from '../../services/offers/actions'
import { bindActionCreators } from 'redux'
import * as offersActions from '../../services/offers/actions'

function OffersPage({ loadOffers }){
    useEffect(()=> {
        loadOffers()
    }, [loadOffers])
    return (
        <div>
            <h2>Ofertas</h2>
        </div>
    )
}

const MapStateToProps = state => ({
    offers: state.offers
})


const MapDispatchToProps = dispatch => 
    bindActionCreators(offersActions, dispatch)

export default connect(MapStateToProps, MapDispatchToProps)(OffersPage)