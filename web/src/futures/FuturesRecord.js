import futuresApi from './api';
import React, { Component } from 'react';
import { Container, Header, Table, Button, Modal, Form, Message } from 'semantic-ui-react';

const defaultFormValues = () => ({
    tiker: '',
    is_open: true,
    warranty_provision: 0,
    margin: 0,
    count: 1,
    commission: 10,
    commission_type: 'fix_price',
});

const commission_type = ['fix_price', 'percent'].map((x) => ({ text: x, value: x }));
const is_open = [
    { value: true, text: 'Да' },
    { value: false, text: 'Нет' },
];

function listFutures() {
    return futuresApi.getFuturesList().then((res) => res.json());
}

function updateFutures(id, params) {
    console.log(id, params);
    return futuresApi.updateFutures(id, params).then((res) => res.json());
}

function createFutures(data) {
    return futuresApi.createFuturesRecord(data).then((res) => res.json());
}

class FuturesList extends Component {
    constructor(props) {
        super(props);

        this.state = {
            futures: [],
            editFutures: null,
            formValues: defaultFormValues(),
            formError: null,
            isModalOpen: false,
        };

        this.submitEdit = this.submitEdit.bind(this);
    }

    componentDidMount() {
        listFutures().then((data) => this.setState({ futures: data.sort((a, b) => (a.id > b.id ? 1 : -1)) }));
    }

    updateEditValue(key, value) {
        this.setState({ formValues: { ...this.state.formValues, [key]: value } });
    }

    submitEdit() {
        const { formValues, editFutures, futures } = this.state;
        formValues.tiker = formValues.tiker.trim();

        if (!formValues.tiker) {
            this.setState({ formError: 'Тикер обязателен для заполнения' });
            return;
        }
        if (!formValues.warranty_provision) {
            this.setState({ formError: 'ГО обязательно для заполнения' });
            return;
        }

        if (editFutures) {
            updateFutures(editFutures.id, {
                margin: Number(formValues.margin),
                is_open: formValues.is_open,
            })
                .then((data) => {
                    for (let i in futures) {
                        if (futures[i].id === data.id) {
                            futures[i] = data;
                            break;
                        }
                    }

                    this.setState({
                        futures,
                        isModalOpen: false,
                        editFutures: null,
                        formValues: defaultFormValues(),
                        formError: null,
                    });
                })
                .catch((error) => {
                    this.setState({
                        formError: `Не удалось сохранить запись: ${error}`,
                    });
                });
        } else {
            createFutures({
                tiker: formValues.tiker,
                is_open: formValues.is_open,
                warranty_provision: Number(formValues.warranty_provision),
                margin: Number(formValues.margin),
                count: Number(formValues.count),
                commission: Number(formValues.commission),
                commission_type: formValues.commission_type,
            })
                .then((data) => {
                    this.setState({
                        futures: this.state.futures.concat(data),
                        isModalOpen: false,
                        editFutures: null,
                        formValues: defaultFormValues(),
                        formError: null,
                    });
                })
                .catch((error) => {
                    this.setState({
                        formError: `Не удалось создать запись: ${error}`,
                    });
                });
        }
    }

    render() {
        const { isModalOpen, editFutures, futures, formValues } = this.state;

        return (
            <Container>
                <Header>
                    Фьючерсы
                    <Button
                        size="tiny"
                        floated="right"
                        color="blue"
                        onClick={() =>
                            this.setState({ editFutures: null, isModalOpen: true, formValues: defaultFormValues() })
                        }
                    >
                        Добавить
                    </Button>
                </Header>

                <Table celled>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>Тикер</Table.HeaderCell>
                            <Table.HeaderCell>ГО</Table.HeaderCell>
                            <Table.HeaderCell>Кол-во</Table.HeaderCell>
                            <Table.HeaderCell>Маржа</Table.HeaderCell>
                            <Table.HeaderCell>Редактировать</Table.HeaderCell>
                        </Table.Row>
                    </Table.Header>

                    <Table.Body>
                        {futures.map((x) => (
                            <Table.Row key={x.id}>
                                <Table.Cell>{x.tiker}</Table.Cell>
                                <Table.Cell>{x.warranty_provision}</Table.Cell>
                                <Table.Cell>{x.count}</Table.Cell>
                                <Table.Cell>{x.margin}</Table.Cell>
                                <Table.Cell>
                                    <Button.Group basic>
                                        <Button
                                            icon="edit"
                                            color="blue"
                                            onClick={() =>
                                                this.setState({
                                                    editFutures: x,
                                                    isModalOpen: true,
                                                    formValues: {
                                                        ...x,
                                                    },
                                                })
                                            }
                                        />
                                    </Button.Group>
                                </Table.Cell>
                            </Table.Row>
                        ))}
                    </Table.Body>
                </Table>

                <Modal open={isModalOpen} onClose={() => this.setState({ isModalOpen: false })}>
                    <Modal.Header>{editFutures ? editFutures.name : 'Новая запись'}</Modal.Header>

                    <Modal.Content>
                        {this.state.formError ? <Message error>{this.state.formError}</Message> : null}
                        <Form
                            onSubmit={(e) => {
                                e.preventDefault();
                                this.submitEdit();
                            }}
                            id="edit-form"
                        >
                            <Form.Group widths="equal">
                                <Form.Input
                                    fluid
                                    label="Тикер"
                                    name="tiker"
                                    value={formValues.tiker}
                                    onChange={(e) => this.updateEditValue('tiker', e.target.value)}
                                />
                            </Form.Group>
                            <Form.Group widths="equal">
                                <Form.Input
                                    fluid
                                    label="Гарантийное обеспечение"
                                    name="warranty_provision"
                                    value={formValues.warranty_provision}
                                    onChange={(e) => this.updateEditValue('warranty_provision', e.target.value)}
                                />
                            </Form.Group>
                            <Form.Group widths="equal">
                                <Form.Input
                                    fluid
                                    label="Вариоационная маржа"
                                    name="margin"
                                    value={formValues.margin}
                                    onChange={(e) => this.updateEditValue('margin', e.target.value)}
                                />
                            </Form.Group>
                            <Form.Group widths="equal">
                                <Form.Select
                                    fluid
                                    label="Позиция открыта"
                                    name="is_open"
                                    options={is_open}
                                    value={formValues.is_open}
                                    // @ts-ignore
                                    onChange={(e, data) => this.updateEditValue('is_open', data.value)}
                                />
                            </Form.Group>
                            <Form.Group widths="equal">
                                <Form.Input
                                    fluid
                                    label="Количество"
                                    name="count"
                                    value={formValues.count}
                                    onChange={(e) => this.updateEditValue('count', e.target.value)}
                                />
                            </Form.Group>
                            <Form.Group widths="equal">
                                <Form.Input
                                    fluid
                                    label="Комиссия"
                                    name="commission"
                                    value={formValues.commission}
                                    onChange={(e) => this.updateEditValue('commission', e.target.value)}
                                />
                            </Form.Group>
                            <Form.Group widths="equal">
                                <Form.Select
                                    fluid
                                    label="Тип комиссии"
                                    name="commission_type"
                                    options={commission_type}
                                    value={formValues.commission_type}
                                    // @ts-ignore
                                    onChange={(e, data) => this.updateEditValue('commission_type', data.value)}
                                />
                            </Form.Group>
                        </Form>
                    </Modal.Content>
                    <Modal.Actions>
                        <Button color="black" onClick={() => this.setState({ isModalOpen: false })}>
                            Close
                        </Button>

                        <Button positive form="edit-form">
                            Сохранить
                        </Button>
                    </Modal.Actions>
                </Modal>
            </Container>
        );
    }
}

export default FuturesList;
