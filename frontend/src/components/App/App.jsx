import React, { Component } from 'react';
import logo from '../../logo.svg';
import './App.css';
import Link from 'react-router-dom';

class App extends Component {
  constructor(props) {
    super(props);
    this.API_URL = process.env.REACT_APP_API_URL;
    this.state = {
      data: "No data",
    };
  }

  componentDidMount() {
    console.log(this.API_URL);
    fetch(this.API_URL + 'test')
      .then(response => response.text())
      .then(data => this.setState({ data }));
  }

  render() {
    let data = this.state.data || "No response";
    console.log(data);
    return (
      <div className="App">
        <ul>
          <li>
            <Link to='/signup'>Signup</Link>
          </li>
        </ul>
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
        <p>
          { data }
        </p>
      </div>
    );
  }
}

export default App;
