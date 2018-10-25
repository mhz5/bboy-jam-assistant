import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './components/App';
import Signup from './components/Signup';
import registerServiceWorker from './registerServiceWorker';
import User from './components/User';
import  { BrowserRouter as Router, Route } from 'react-router-dom';
import Grid from '@material-ui/core/Grid';
import Login from "./components/Login/Login";
import { createMuiTheme, MuiThemeProvider } from '@material-ui/core/styles';
import blue from '@material-ui/core/colors/blue';

const theme = createMuiTheme({
  palette: {
    primary: blue,
    secondary: {
      main: '#f44336',
    },
  },
});
ReactDOM.render(
  <Router>
    <MuiThemeProvider theme={theme}>
    <div className="App-wrapper">
        <Route path="/" component={App} />
        <div className="app-content-wrapper">
          <Route path="/signup" component={Signup} />
          <Route path="/login" component={Login} />
          <Route path="/users/:userId" component={User} />
        </div>
    </div>
    </MuiThemeProvider>
  </Router>,
  document.getElementById('root')
);
registerServiceWorker();
