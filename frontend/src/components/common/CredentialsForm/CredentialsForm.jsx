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


class CredentialsForm extends Component {
  constructor(props) {
    super(props);
    this.submitActionUrl = props.submitActionUrl;
    this.action = props.action;

    this.state={
      username:"",
      password: "",
      isPasswordVisible: false
    }
  };

  static contextTypes = {
    router: PropTypes.object
  };

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
    const FD = new FormData();
    FD.append('username', username);
    FD.append('password', password);
    fetch(this.submitActionUrl, {
      method: 'POST',
      body: FD,
      credentials: 'include',
    }).then((res) => {
        return res.json();
    }).then(data => {
        // TODO: this.context is deprecated.
        const redicrectPath = this.props.redirectPath || this.props.getRedirectPath(data);
        this.context.router.history.push(redicrectPath);
    }).catch((e) => {
      console.log('request error:', e);
    });;
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
