import './assets/fomantic/dist/semantic.css';
import React, { Component } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import { Container } from 'semantic-ui-react';

import SideBar from './SideBar';
import Main from './Main';
import PassGen from './pass-gen/Pass-gen';

class App extends Component {
    render() {
        return (
            <BrowserRouter>
                <SideBar />
                <Container>
                    <Switch>
                        <Route exact path="/pass-gen" component={PassGen} />
                        <Route path="/" component={Main} />
                    </Switch>
                </Container>
            </BrowserRouter>
        );
    }
}

export default App;
