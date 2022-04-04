import './assets/fomantic/dist/semantic.css';
import React, { Component } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import { Container } from 'semantic-ui-react';

import SideBar from './SideBar';
import Main from './Main';
import PassGen from './pass-gen/Pass-gen';
import TradingRecord from './trading/TradingRecord';
import TradingLog from './trading/TradingLog';
import FuturesList from './futures/FuturesRecord';

class App extends Component {
    render() {
        return (
            <BrowserRouter>
                <SideBar />
                <Container>
                    <Switch>
                        <Route exact path="/pass-gen" component={PassGen} />
                        <Route exact path="/trading" component={TradingRecord} />
                        <Route exact path="/trading/log" component={TradingLog} />
                        <Route exact path="/futures" component={FuturesList} />
                        <Route path="/" component={Main} />
                    </Switch>
                </Container>
            </BrowserRouter>
        );
    }
}

export default App;
