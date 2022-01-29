import { Container, Form, Header, Input, Message, Segment } from 'semantic-ui-react';
import { Component } from 'react';

import passGenApi from './api';


class PassGen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            context: props.context,
            error: null,
            isPassLoaded: false,
            isPassRequested: false,
            pass: null,
            lowerCaseCount: null,
            upperCaseCount: null,
            numbersCount: null,
            specialChairCount: null
        };

        this.getPass = this.getPass.bind(this);
        this.handleChangeValue = this.handleChangeValue.bind(this);
    }

    handleChangeValue(field, value) {
        this.setState({ ...this.state, [field]: value });
    }

    async getPass() {
        const { lowerCaseCount, upperCaseCount, numbersCount, specialChairCount } = this.state;
        this.setState({
            error: null,
            isPassLoaded: false,
            isPassRequested: true,
            pass: null
        });

        if (!lowerCaseCount || parseInt(lowerCaseCount) < 3) {
            this.setState({
                error: { message: 'Для надежного пароля кол-во символов в нижнем регистре должно быть больше 3' },
                isPassRequested: false,
            });
            return;
        }

        if (!upperCaseCount || parseInt(upperCaseCount) < 3) {
            this.setState({
                error: { message: 'Для надежного пароля кол-во символов в верхнем регистре должно быть больше 3' },
                isPassRequested: false,
            });
            return;
        }

        if (!numbersCount || parseInt(numbersCount) < 2) {
            this.setState({
                error: { message: 'Для надежного пароля кол-во цифр должно быть больше 2' },
                isPassRequested: false,
            });
            return;
        }

        if (!specialChairCount || parseInt(specialChairCount) < 2) {
            this.setState({
                error: { message: 'Для надежного пароля кол-во специальных символов должно быть больше 2' },
                isPassRequested: false,
            });
            return;
        }

        const params = { lowerCase: lowerCaseCount, upperCase: upperCaseCount, numbers: numbersCount, specialChair: specialChairCount };

        const resp = await passGenApi.getPass(params);

        if (resp.ok) {
            const result = await resp.json();
            this.setState({
                pass: result,
                isPassLoaded: true,
                error: null,
                lowerCaseCount: null,
                upperCaseCount: null,
                numbersCount: null,
                specialChairCount: null
            });
        } else {
            console.error(resp)
            this.setState({
                error: { message: await resp.text() },
                isPassRequested: false,
                lowerCaseCount: null,
                upperCaseCount: null,
                numbersCount: null,
                specialChairCount: null
            });
        }
    }

    render() {
        const {
            error,
            isPassLoaded,
            isPassRequested,
            pass,
        } = this.state;
        return (
            <Container>
                <Header>Сгенерировать пароль</Header>
                {error && (
                    <Message negative>
                        <Message.Header>Произошла ошибка</Message.Header>
                        <p>{error.message}</p>
                    </Message>
                )}
                <Segment>
                    <Form onSubmit={this.getPass}>
                        <Form.Group widths="equal">
                            <Form.Field inline>
                                <label>Кол-во символов в нижнем регистре</label>
                                <Input
                                    placeholder="5"
                                    name="lowerCaseCount"
                                    onChange={(e) => this.handleChangeValue('lowerCaseCount', e.target.value)}
                                />
                            </Form.Field>
                            <Form.Field inline>
                                <label>Кол-во символов в верхнем регистре</label>
                                <Input
                                    placeholder="5"
                                    name="upperCaseCount"
                                    onChange={(e) => this.handleChangeValue('upperCaseCount', e.target.value)}
                                />
                            </Form.Field>
                            <Form.Field inline>
                                <label>Кол-во цифр</label>
                                <Input
                                    placeholder="5"
                                    name="numbersCount"
                                    onChange={(e) => this.handleChangeValue('numbersCount', e.target.value)}
                                />
                            </Form.Field>
                            <Form.Field inline>
                                <label>Кол-во специальных символов</label>
                                <Input
                                    placeholder="5"
                                    name="specialChairCount"
                                    onChange={(e) => this.handleChangeValue('specialChairCount', e.target.value)}
                                />
                            </Form.Field>
                        </Form.Group>
                        <Form.Group>
                            <Form.Button fluid primary>Run</Form.Button>
                        </Form.Group>
                    </Form>
                </Segment>
                {isPassRequested &&
                    <Segment loading={!isPassLoaded}>
                        <Header as='h3'>{pass}</Header>
                    </Segment>
                }
            </Container>
        );
    }
}

export default PassGen;
