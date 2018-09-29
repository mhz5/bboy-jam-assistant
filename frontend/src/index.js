import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './components/App';
import Signup from './components/Signup';
import registerServiceWorker from './registerServiceWorker';
import  { BrowserRouter as Router, Route } from 'react-router-dom';

ReactDOM.render(
  //<App />
  <Router>
    <div>
      <Route exact path="/" component={App} />
      <Route path="/signup" component={Signup} />
    </div>
  </Router>,
  document.getElementById('root')
);
registerServiceWorker();
