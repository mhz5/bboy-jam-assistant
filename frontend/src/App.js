import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
    constructor(props) {
        super(props);

        this.state = {
            data: null,
        };
    }

    componentDidMount() {
        fetch('https://backend-dot-bboy-jam-prod.appspot.com/test')
            .then(response => response.text())
            .then(data => this.setState({ data }));
    }

    render() {
        let data = this.state.data;
        return (
            <div className="App">
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
