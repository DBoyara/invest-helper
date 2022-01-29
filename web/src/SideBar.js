import React, { Component } from 'react';
import { Menu, Image } from 'semantic-ui-react';
import { NavLink } from 'react-router-dom';

class SideBar extends Component {
    render() {
        return (
            <Menu>
                <NavLink exact component={Menu.Item} to="/">
                    <Image size="mini" src="/logo192.png" />
                </NavLink>

                <Menu.Menu position="left">
                    <NavLink component={Menu.Item} to="/pass-gen">
                        Генератор паролей
                    </NavLink>
                </Menu.Menu>
            </Menu>
        );
    }
}

export default SideBar;
