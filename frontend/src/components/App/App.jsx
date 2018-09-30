import React, { Component } from 'react';
import logo from '../../logo.svg';
import './App.scss';

class App extends Component {
  // constructor(props) {
  //   super(props);
  //   this.API_URL = process.env.REACT_APP_API_URL;
  //   this.state = {
  //     data: "No data",
  //   };
  // }
  //
  // componentDidMount() {
  //   console.log(this.API_URL);
  //   fetch(this.API_URL + 'test')
  //     .then(response => response.text())
  //     .then(data => this.setState({ data }));
  // }

  render() {
    return (
      <div className="home">
        <h2>B-BOY DANCE BATTLE HOME</h2>
      </div>
    );
  }
}

export default App;
