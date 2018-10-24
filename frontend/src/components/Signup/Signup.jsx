import React, { Component } from 'react';
import './Signup.scss';
import CredentialsForm from 'components/common/CredentialsForm';


class Signup extends Component {
  getRedirectPath(res) {
    res.json().then((resp) => {
      console.log('resp', resp);
    })

  };
  render() {
    return (
      <CredentialsForm
        action="Signup"
        getRedirectPath={this.getRedirectPath}
        submitActionUrl="http://192.168.1.142:8080/users"
      />
    );
  }
}

export default Signup;
