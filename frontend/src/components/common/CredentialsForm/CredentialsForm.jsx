import React, { Component } from 'react';
import './CredentialsForm.scss';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import InputAdornment from '@material-ui/core/InputAdornment';
import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';
import IconButton from '@material-ui/core/IconButton';
import Button from '@material-ui/core/Button';
import PropTypes from 'prop-types';

const protobuf = require('protobufjs');


class CredentialsForm extends Component {
  constructor(props) {
    super(props);
    this.submitActionUrl = props.submitActionUrl;
    this.action = props.action;
    console.log(protobuf);

    // TODO: There MUST be a better way of storing a reference to type AuthRequest.
    this.AuthRequest = null;
    protobuf.load('protos/sixstep_request.proto', (err, root) => {
      if (err)
        throw err;
      this.AuthRequest = root.lookupType('sixstep.AuthRequest');
    });

    this.state={
      username:"",
      password: "",
      isPasswordVisible: false
    }
  };

  static contextTypes = {
    router: PropTypes.object
  }

  onInputChange = type => e => {
    this.setState({
      [type]: e.target.value
    });
  };

  toggleShowPassword = e => {
    this.setState(state => ({ isPasswordVisible: !state.isPasswordVisible }));
  };

  onSubmit = () => {
    const { username, password } = this.state;

    let payload = {
      username: username,
      password: password,
    };

    let err = this.AuthRequest.verify(payload);
    if (err)
      throw Error(err);
    let message = this.AuthRequest.create(payload);
    let buffer = this.AuthRequest.encode(message).finish();

    fetch(this.submitActionUrl, {
      method: 'POST',
      body: buffer,
      credentials: 'include',
    }).then((res) => {
      if (res.ok) {
        this.context.router.history.push(this.props.redirectPath);
      }
      });
  };

  render() {
    return (
      <div className="credentials-form__wrapper">
        <h2>{this.action}</h2>
        {/* Username input */}
        <InputLabel className="credentials-form__label"
          htmlFor="username-input">
          Username:
        </InputLabel>
        <Input
          className="credentials-form__input credentials-form__input--username"
          id="username-input"
          value={this.state.username}
          onChange={this.onInputChange("username")} />

        {/* Password input */}
        <InputLabel
          className="credentials-form__label"
          htmlFor="password-input">
          Password:
        </InputLabel>
        <Input
          className="credentials-form__input credentials-form__input--password"
          id="password-input"
          type={this.state.isPasswordVisible ? "text" : "password"}
          value={this.state.password}
          onChange={this.onInputChange("password")}
          endAdornment={
            <InputAdornment position="end">
              <IconButton
                aria-label="Toggle password visibility"
                onClick={this.toggleShowPassword}>
                {this.state.isPasswordVisible ? <VisibilityOff/> : <Visibility/>}
              </IconButton>
            </InputAdornment>
          }/>

        {/* Submit CTA */}
        <Button
          variant="contained"
          color="primary"
          onClick={this.onSubmit}>{this.action}</Button>
      </div>
    );
  }
}

export default CredentialsForm;
