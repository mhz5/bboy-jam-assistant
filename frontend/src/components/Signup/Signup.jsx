import React, { Component } from 'react';
import './Signup.scss';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import InputAdornment from '@material-ui/core/InputAdornment';
import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';
import IconButton from '@material-ui/core/IconButton';
import Button from '@material-ui/core/Button';

import FormControl from '@material-ui/core/FormControl';
import { withStyles } from '@material-ui/core/styles';


class Signup extends Component {
  constructor(props) {
    super(props);
    this.state={
      username:"",
      password: "",
      isShowPassword: false
    }
  };

  onInputChange = type => e => {
    this.setState({
      [type]: e.target.value
    });
  };

  toggleShowPassword = e => {
    this.setState(state => ({ isShowPassword: !state.isShowPassword }));
  };

  onClickSignup = () => {
    const url = "http://localhost:8080/users";
    const { username, password } = this.state;
    const FD = new FormData();
    FD.append('username', username);
    FD.append('password', password);
    fetch(url, {
      method: 'POST',
      body: FD,
    })
      .then(res => res.text())
      .then(data => {
      console.log('res::', JSON.parse(data));
    }).catch();
  };
  render() {
    return (
      <div className="sign-up-form__wrapper">
        <InputLabel
          className="sign-up-form__label"
          htmlFor="username-input">
          Username:
        </InputLabel>
        <Input
          className="sign-up-form__input sign-up-form__input--username"
          id="username-input"
          value={this.state.username}
          onChange={this.onInputChange("username")} />


        <InputLabel
          className="sign-up-form__label"
          htmlFor="password-input">
          Password:
        </InputLabel>
        <Input
          className="sign-up-form__input sign-up-form__input--password"
          id="password-inpu"
          type={this.state.isShowPassword ? "text" : "password"}
          value={this.state.password}
          onChange={this.onInputChange("password")}
          endAdornment={
            <InputAdornment position="end">
              <IconButton
                aria-label="Toggle password visibility"
                onClick={this.toggleShowPassword}>
                {this.state.isShowPassword ? <VisibilityOff/> : <Visibility/>}
              </IconButton>
            </InputAdornment>
          }/>
        <Button
          variant="contained"
          color="primary"
          onClick={this.onClickSignup}> Sign Up </Button>
      </div>
    );
  }
}

export default Signup;
