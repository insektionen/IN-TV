import {Card, Col} from 'react-bootstrap';
import {ScreenPlayer} from './ScreenPlayer';

export const FakeScreen = (props: { name: string }) => {
    return (
        <Col>
            <Card>
                <Card.Header>{props.name}</Card.Header>
                <Card.Body>
                    <ScreenPlayer name={props.name} previews/>
                </Card.Body>
            </Card>
        </Col>
    );
};