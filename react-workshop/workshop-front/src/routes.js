import React from 'react'
import { BrowserRouter, Route } from 'react-router-dom'
import Login from './pages/Login/index'
import Offer from './pages/Offers/index'

function Routes() {
    return (
        <BrowserRouter>
            <Route exac path="/" component={Login}></Route>
            <Route path="/ofertas" component={Offer}></Route>
        </BrowserRouter>
    )
}

export default Routes