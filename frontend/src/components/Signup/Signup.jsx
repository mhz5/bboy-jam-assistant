import React, { Component } from 'react';
import './Signup.scss';
import CredentialsForm from 'components/common/CredentialsForm';


class Signup extends Component {
  getRedirectPath(data) {
    return `/users/${data.id}`
  };
  render() {
    return (
      <CredentialsForm
        action="Signup"
        getRedirectPath={this.getRedirectPath}
        submitActionUrl={`${process.env.REACT_APP_API_URL}/users`}
      />
    );
  }
}

export default Signup;
