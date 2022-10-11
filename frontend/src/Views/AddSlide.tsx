import {PictureThumbnail} from '../Components/Slides/Picture/PictureThumbnail';
import {Slide, Slideshow} from '../API/Types';
import {TextThumbnail} from '../Components/Slides/Text/TextThumbnail';
import {Row} from 'react-bootstrap';
import {ClockThumbnail} from '../Components/Slides/Clock/ClockThumbnail';
import React, {useContext, useEffect, useState} from 'react';
import {useLocation, useNavigate} from 'react-router-dom';
import {ApiContext} from '../App';
import {EditSlideModal} from '../Components/Modal/EditSlideModal';
import {AlertContext} from '../Components/Alert';

const AllThumbnails = () => {
    const emptySlide: Slide = {type: 'empty', position: 0, timeout: 0, data: {}};
    const [addModalVisible, setAddModalVisible] = useState(false);
    const [selectedType, setSelectedType] = useState('');
    const [currentSlideshow, setCurrentSlideshow] = useState<Slideshow>({name: '', slides: []});
    const cli = useContext(ApiContext);
    const {setData: setAlertData} = useContext(AlertContext);
    const location = useLocation();
    const navigate = useNavigate();

    const onClick = (s: Slide) => {
        setSelectedType(s.type);
        setAddModalVisible(true);
    };
    const onClose = () => {
        setSelectedType('empty');
        setAddModalVisible(false);
    };

    const onSave = (s: Slide) => {
        const toSend = currentSlideshow;
        s.position = toSend.slides.length;
        toSend.slides.push(s);
        cli.updateSlideshow(toSend).then(sl => {
            setCurrentSlideshow(sl);
            setAlertData({text: 'Added!', timeout: 2, variant: 'success'});
            navigate('/edit', {state: {selectedSlideshow: sl}});
        });
    };

    useEffect(() => {
        if (location.state) {
            const data = location.state as { slideshow: Slideshow };
            setCurrentSlideshow(data.slideshow);
        }
    }, [location.state]);

    return (
        <>
            <Row xs={1} md={2} className={'g-4'}>
                <PictureThumbnail slide={{...emptySlide, type: 'picture'}} onClick={onClick} />
                <TextThumbnail slide={{...emptySlide, type: 'text'}} onClick={onClick} />
                <ClockThumbnail slide={{...emptySlide, type: 'clock'}} onClick={onClick} />
            </Row>
            <EditSlideModal visible={addModalVisible} onClose={onClose} okButtonText={'Add'}
                slide={{...emptySlide, type: selectedType}} onSave={onSave} />
        </>
    );
};

export const AddSlide = () => {
    return (
        <div>
            <h2>Add new slide:</h2>
            <AllThumbnails />
        </div>
    );
};