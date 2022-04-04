import React, { Component } from 'react';
import { Menu, Image, Dropdown } from 'semantic-ui-react';
import { NavLink } from 'react-router-dom';

class SideBar extends Component {
    render() {
        return (
            <Menu>
                <NavLink exact component={Menu.Item} to="/">
                    <Image size="mini" src="/logo192.png" />
                </NavLink>

                <Dropdown item text="Пароли">
                    <Dropdown.Menu>
                        <NavLink component={Dropdown.Item} to="/pass-gen">
                            Генератор паролей
                        </NavLink>
                    </Dropdown.Menu>
                </Dropdown>

                <Dropdown item text="Trading">
                    <Dropdown.Menu>
                        <NavLink component={Menu.Item} to="/trading">
                            TradingRecord
                        </NavLink>
                        <NavLink component={Menu.Item} to="/trading/log">
                            TradingLog
                        </NavLink>
                    </Dropdown.Menu>
                </Dropdown>

                <Dropdown item text="Фьючерсы">
                    <Dropdown.Menu>
                        <NavLink component={Menu.Item} to="/futures">
                            Фьючерсы
                        </NavLink>
                    </Dropdown.Menu>
                </Dropdown>
            </Menu>
        );
    }
}

export default SideBar;
