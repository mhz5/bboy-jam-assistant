import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './components/App';
import Signup from './components/Signup';
import registerServiceWorker from './registerServiceWorker';
import  { BrowserRouter as Router, Route } from 'react-router-dom';
import Grid from '@material-ui/core/Grid';

ReactDOM.render(
  //<App />
  <Router>
    <div className="App-wrapper">
      <Grid container spacing={24}>
        <Grid item xs={12} justify="center">
          <Route path="/" component={App} />
          <Route path="/signup" component={Signup} />
        </Grid>
      </Grid>
    </div>
  </Router>,
  document.getElementById('root')
);
registerServiceWorker();
