import React from 'react'
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";
import Navbar from './components/navbar/Navbar';
import NewProduct from './components/product/NewProduct';
import ProductsList from './components/product/ProductsList';
import Login from './components/user/Login';
import Register from './components/user/Register';
import WhoAmI from './components/user/WhoAmI';

function App() {
  return (
    <div className="App">
      <Router>
        <Navbar />
        <Switch>   
          <Route path="/" exact>
            <ProductsList />
          </Route>
          <Route path="/register" exact>
            <Register />
          </Route>
          <Route path="/login" exact>
            <Login />
          </Route>
           <Route path="/whoami" exact>
            <WhoAmI />
          </Route>
          <Route path="/product" exact>
            <NewProduct />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
