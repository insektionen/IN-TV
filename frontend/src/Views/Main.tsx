import {Button, Col, Container, ListGroup, Row} from 'react-bootstrap';
import React, {ReactChild, useContext, useEffect, useState} from 'react';
import {Status} from '../API/Types';
import {ApiContext} from '../App';
import {Loading} from '../Components/Loading';
import {FakeScreen} from '../Components/FakeScreen';
import {useNavigate} from 'react-router-dom';

const FlexCol = (props: { children: ReactChild | ReactChild[] }) => {
    return (
        <Col style={{display: 'flex', flexDirection: 'column'}}>{props.children}</Col>
    );
};

const CurrentlyRunning = (props: { name: string, onStop: (name: string) => void }) => {
    return (
        <div style={{display: 'flex', alignItems: 'center'}}>
            <div style={{width: '100%'}}>{props.name}</div>
            <Button style={{float: 'right'}} variant={'danger'} onClick={() => props.onStop(props.name)}>Stop</Button>
        </div>
    );
};

export const Main = () => {
    const [currentStatus, setCurrentStatus] = useState<Status>();
    const navigate = useNavigate();
    const cli = useContext(ApiContext);

    useEffect(() => {
        cli.getStatus().then(res => setCurrentStatus(res));
    }, []);

    if (!currentStatus) {
        return <Loading />;
    }

    const handleStartButton = () => {
        navigate('start');
    };
    const onStopClick = (name: string) => {
        cli.stopSlideshow(name).then(() => {
            cli.getStatus().then(res => setCurrentStatus(res));
        });
    };

    return (
        <Container>
            <Row style={{display: 'flex', flexWrap: 'wrap'}}>
                <FlexCol>
                    <Row>
                        <h2>Currently running:</h2>
                    </Row>
                    <div style={{
                        display: 'flex',
                        flexDirection: 'column',
                        justifyContent: 'space-between',
                        height: '100%'
                    }}>
                        <Row>
                            {currentStatus.running_slideshows.length === 0 ?
                                <h4>Currently no slideshows running</h4> :
                                <ListGroup>
                                    {
                                        currentStatus.running_slideshows.map(name => (
                                            <ListGroup.Item key={name}><CurrentlyRunning name={name}
                                                onStop={onStopClick} /></ListGroup.Item>))
                                    }
                                </ListGroup>}
                        </Row>
                        <Row>
                            <Col>
                                <Button variant="primary" onClick={handleStartButton}>Start
                                    slideshow</Button>
                            </Col>
                        </Row>
                    </div>
                </FlexCol>
                <FlexCol>
                    <Row>
                        <h2>Preview</h2>
                    </Row>
                    <Row xs={1} md={2} className={'g-2'}>
                        {
                            currentStatus.connected_screens.map(screen => <FakeScreen name={screen.name}
                                key={screen.name} />)
                        }
                    </Row>
                </FlexCol>
            </Row>
        </Container>
    );
};