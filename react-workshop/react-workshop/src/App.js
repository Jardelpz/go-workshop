import React from 'react';
import './App.css'
import Header from './Header'
import Index from './button/Index'
import { Provider } from 'react-redux'
import store from './store'
import Routes from './routes'
//const app = () => {}
function App() {
  return (
    // <div className="App">
    //   Hello World!
    //   <Header text="Olá" value="Jardeel"/>
    //   <Index title="Este é um títlo" cor="#F42"/>
    //   {/* <Conteudo title="Este é um títlo" pagina={2}/> */}
    // </div>
    <Provider store={store}>
      <Routes/>
    </Provider>
  )
}

export default App;
