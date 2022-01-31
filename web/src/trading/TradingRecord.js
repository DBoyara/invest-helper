import { Container, Form, Header, Input, Message, Segment } from 'semantic-ui-react';
import { Component } from 'react';

import tradingApi from './api';

const tradeTypeOptions = [
    { value: 'sell', text: 'Продажа', key: 1 },
    { value: 'buy', text: 'Покупка', key: 2 },
];

class TradingRecord extends Component {
    constructor(props) {
        super(props);
        this.state = {
            context: props.context,
            error: null,
            isCommissionsLoaded: false,
            isRecordCreated: false,
            tiker: null,
            price: null,
            count: null,
            lot: null,
            commissionsOptions: [],
            commissionSelected: [],
            tradeTypeSelected: [],
        };

        this.createRecord = this.createRecord.bind(this);
        this.handleChangeValue = this.handleChangeValue.bind(this);
    }

    componentDidMount() {
        this.getCommissions();
    }

    handleChangeValue(field, value) {
        this.setState({ ...this.state, [field]: value });
    }

    async getCommissions() {
        const resp = await tradingApi.getCommissions();

        if (resp.ok) {
            const result = await resp.json();
            this.setState({
                data: result,
                isCommissionsLoaded: true,
                error: null,
                commissionsOptions: result.map((data, key) => {
                    return {
                        value: data.value,
                        text: data.value,
                        type: data.type,
                        key: key,
                    };
                }),
            });
        } else {
            console.error(resp);
            this.setState({
                error: { message: resp },
                isCommissionsLoaded: true,
            });
        }
    }

    async createRecord() {
        const { commissionSelected, commissionsOptions, tiker, price, tradeTypeSelected, count, lot } = this.state;
        let commissionType;
        for (let commissions of commissionsOptions) {
            if (commissionSelected === commissions.value) {
                commissionType = commissions.type;
            }
        }

        const curPrice = parseFloat(price);
        const curCount = parseFloat(count);
        const curLot = parseFloat(lot);
        const curTiker = tiker.toUpperCase();
        const params = {
            tiker: curTiker,
            type: tradeTypeSelected,
            price: curPrice,
            count: curCount,
            lot: curLot,
            commission: commissionSelected,
            commission_type: commissionType,
        };

        const resp = await tradingApi.createTradeRecord(params);

        if (resp.ok) {
            const result = await resp.json();
            console.log(result);
            this.setState({
                isRecordCreated: true,
                error: null,
            });
        } else {
            console.error(resp);
            this.setState({
                error: { message: await resp.text() },
            });
        }
    }

    async getRecords() {}

    render() {
        const { error, isCommissionsLoaded, commissionsOptions, isRecordCreated } = this.state;
        return (
            <Container>
                <Header>Запись о сделке</Header>
                {error && (
                    <Message negative>
                        <Message.Header>Произошла ошибка</Message.Header>
                        <p>{error.message}</p>
                    </Message>
                )}
                {isRecordCreated && (
                    <Message success>
                        <Message.Header>Успешно!</Message.Header>
                    </Message>
                )}
                <Segment loading={!isCommissionsLoaded}>
                    <Form onSubmit={this.createRecord}>
                        <Form.Group widths="equal">
                            <Form.Field>
                                <label>Тикер</label>
                                <Input
                                    required
                                    placeholder="SBER"
                                    name="tiker"
                                    onChange={(e) => this.handleChangeValue('tiker', e.target.value)}
                                />
                            </Form.Field>
                            <Form.Field>
                                <label>Цена</label>
                                <Input
                                    required
                                    placeholder="100"
                                    name="price"
                                    onChange={(e) => this.handleChangeValue('price', e.target.value)}
                                />
                            </Form.Field>
                            <Form.Field>
                                <label>Кол-во</label>
                                <Input
                                    required
                                    placeholder="1"
                                    name="count"
                                    onChange={(e) => this.handleChangeValue('count', e.target.value)}
                                />
                            </Form.Field>
                            <Form.Field>
                                <label>Лотность</label>
                                <Input
                                    required
                                    placeholder="100"
                                    name="lot"
                                    onChange={(e) => this.handleChangeValue('lot', e.target.value)}
                                />
                            </Form.Field>
                        </Form.Group>
                        <Form.Group widths="equal">
                            <Form.Select
                                inline
                                fluid
                                label="Комиссия"
                                placeholder="Комиссия"
                                clearable
                                search
                                required
                                options={commissionsOptions}
                                name="commissionSelected"
                                onChange={(e, data) => this.handleChangeValue('commissionSelected', data.value)}
                            />
                            <Form.Select
                                inline
                                fluid
                                label="Тип сделки"
                                clearable
                                search
                                required
                                options={tradeTypeOptions}
                                name="tradeTypeSelected"
                                onChange={(e, data) => this.handleChangeValue('tradeTypeSelected', data.value)}
                            />
                        </Form.Group>
                        <Form.Group>
                            <Form.Button fluid>Записать сделку</Form.Button>
                        </Form.Group>
                    </Form>
                </Segment>
            </Container>
        );
    }
}

export default TradingRecord;
