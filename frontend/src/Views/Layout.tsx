import {Button, Container, Dropdown, Form, Modal, Nav, Navbar, NavDropdown} from 'react-bootstrap';
import {SlideshowList} from '../Components/SlideshowList';
import {Outlet, useNavigate} from 'react-router-dom';
import React, {useContext, useState} from 'react';
import {ApiContext} from '../App';
import {Slideshow} from '../API/Types';
import {Alert, AlertContext} from '../Components/Alert';

export const Layout = () => {
    const [addModalVisible, setAddModalVisible] = useState(false);
    const [addName, setAddName] = useState('');
    const {setData: setAlertData} = useContext(AlertContext);
    const navigate = useNavigate();
    const cli = useContext(ApiContext);

    const handleAddCancel = () => setAddModalVisible(false);
    const handleAddOk = () => {
        setAddModalVisible(false);
        cli.createSlideshow({name: addName, slides: []}).then(res => {
            setAlertData({variant: 'success', timeout: 2, text: 'New slideshow created'});
            navigate('/edit', {state: {selectedSlideshow: res}});
        });
    };


    return (
        <>
            <div>
                <Navbar bg={'dark'} variant={'dark'} expand={'lg'}>
                    <Container fluid>
                        <Navbar.Brand href={'/'}>IN-TV Admin</Navbar.Brand>
                        <Navbar.Toggle aria-controls={'basic-navbar-nav'} />
                        <Navbar.Collapse id={'navbar-nav'}>
                            <Nav>
                                <Nav.Link href={'/'}>Start</Nav.Link>
                                <NavDropdown title={'Slideshow'} id={'nav-dropdown'}>
                                    <Dropdown.Item onClick={() => setAddModalVisible(true)}>Add new</Dropdown.Item>
                                    <NavDropdown.Divider />
                                    <SlideshowList />
                                </NavDropdown>
                            </Nav>
                        </Navbar.Collapse>
                    </Container>
                </Navbar>
                <div style={{height: '100%', width: '100%', position: 'absolute', marginTop: '15px'}}>
                    <Container>
                        <Alert />
                        <Outlet />
                    </Container>
                </div>
            </div>
            <Modal show={addModalVisible} onHide={handleAddCancel}>
                <Modal.Header closeButton>
                    <Modal.Title>Add new slideshow</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form.Group>
                        <Form.Label>Name</Form.Label>
                        <Form.Control type="text" autoFocus onChange={(e) => setAddName(e.target.value)} />
                    </Form.Group>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleAddCancel}>Close</Button>
                    <Button variant="primary" onClick={handleAddOk}>Add</Button>
                </Modal.Footer>
            </Modal>
        </>
    );
};
