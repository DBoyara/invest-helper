import { Container, Icon, Header, Table, Message, Segment } from 'semantic-ui-react';
import { Component } from 'react';

import tradingApi from './api';

export class StatusCell extends Component {
    render() {
        if (this.props.isOpen) {
            return (
                <Table.Cell textAlign="center">
                    <Icon name="checkmark" color="green" size="large" />
                </Table.Cell>
            );
        }
        if (!this.props.isOpen) {
            return (
                <Table.Cell textAlign="center">
                    <Icon name="attention" color="red" size="large" />
                </Table.Cell>
            );
        }
        return (
            <Table.Cell textAlign="center">
                <Icon name="circle outline" color="grey" size="large" />
            </Table.Cell>
        );
    }
}

export class TypeCell extends Component {
    render() {
        if (this.props.typeLog === 'sell') {
            return (
                <Table.Cell>
                    <Icon name="angle double down" color="red" size="large" />
                </Table.Cell>
            );
        }
        return (
            <Table.Cell>
                <Icon name="angle double up" color="green" size="large" />
            </Table.Cell>
        );
    }
}

class ContactsTable extends Component {
    render() {
        const rows = this.props.logs.map((data, key) => {
            let date = new Date(data.datetime).toLocaleString();
            return (
                <Table.Row key={key}>
                    <Table.Cell>{date}</Table.Cell>
                    <Table.Cell>{data.tiker}</Table.Cell>
                    <TypeCell typeLog={data.type} />
                    <StatusCell isOpen={data.is_open} />
                    <Table.Cell>{data.price}</Table.Cell>
                    <Table.Cell>{data.count}</Table.Cell>
                    <Table.Cell>{data.lot}</Table.Cell>
                    <Table.Cell>{data.amount}</Table.Cell>
                    <Table.Cell>{data.commission_amount}</Table.Cell>
                </Table.Row>
            );
        });

        if (rows.length > 0) {
            return (
                <Table celled padded>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>Время</Table.HeaderCell>
                            <Table.HeaderCell>Тикер</Table.HeaderCell>
                            <Table.HeaderCell>Тип</Table.HeaderCell>
                            <Table.HeaderCell>Завершено</Table.HeaderCell>
                            <Table.HeaderCell>Цена</Table.HeaderCell>
                            <Table.HeaderCell>Кол-во</Table.HeaderCell>
                            <Table.HeaderCell>Лотность</Table.HeaderCell>
                            <Table.HeaderCell>Сумма</Table.HeaderCell>
                            <Table.HeaderCell>Комиссия</Table.HeaderCell>
                        </Table.Row>
                    </Table.Header>
                    <Table.Body>{rows}</Table.Body>
                </Table>
            );
        } else {
            return <Message>Нет записей</Message>;
        }
    }
}

class TradingLog extends Component {
    constructor(props) {
        super(props);
        this.state = {
            context: props.context,
            error: null,
            isLogsLoaded: false,
            logs: [],
        };

        this.handleChangeValue = this.handleChangeValue.bind(this);
    }

    componentDidMount() {
        this.getRecords();
    }

    handleChangeValue(field, value) {
        this.setState({ ...this.state, [field]: value });
    }

    async getRecords() {
        const resp = await tradingApi.getTradeRecords();

        if (resp.ok) {
            const result = await resp.json();
            this.setState({
                logs: result,
                isLogsLoaded: true,
                error: null,
            });
        } else {
            console.error(resp);
            this.setState({
                error: { message: resp },
                isLogsLoaded: true,
            });
        }
    }

    render() {
        const { error, isLogsLoaded, logs } = this.state;
        return (
            <Container>
                <Header>Запись о сделке</Header>
                {error && (
                    <Message negative>
                        <Message.Header>Произошла ошибка</Message.Header>
                        <p>{error.message}</p>
                    </Message>
                )}
                <Segment loading={!isLogsLoaded}>
                    <ContactsTable logs={logs} />
                </Segment>
            </Container>
        );
    }
}

export default TradingLog;
