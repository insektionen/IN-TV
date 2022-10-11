import {Button, Row} from 'react-bootstrap';
import React, {useContext, useEffect, useState} from 'react';
import {useLocation, useNavigate} from 'react-router-dom';
import {Slide, Slideshow} from '../API/Types';
import {SlideThumbnail} from '../Components/Slides/SlideThumbnail';
import {ApiContext} from '../App';
import {EditSlideModal} from '../Components/Modal/EditSlideModal';
import {AlertContext} from '../Components/Alert';

export const Edit = (props: { selectedSlideshow: Slideshow }) => {
    const [editingSlide, setEditingSlide] = useState<Slide>();
    const [selectedSlideshow, setSelectedSlideshow] = useState<Slideshow>(props.selectedSlideshow);
    const {setData: setAlertData} = useContext(AlertContext);
    const cli = useContext(ApiContext);
    const location = useLocation();
    const navigate = useNavigate();

    const handleAddSlide = () => {
        navigate('add_slide', {state: {slideshow: selectedSlideshow}});
    };
    const handleRemoveSlide = (slide: Slide) => {
        const sendData = selectedSlideshow;
        const index = sendData.slides.findIndex(s => s.position === slide.position);
        sendData.slides.splice(index, 1);
        cli.updateSlideshow(sendData).then(sl => {
            setSelectedSlideshow(sl);
            setAlertData({text: 'Removed', timeout: 2, variant: 'success'});
        });
    };
    const handleSaveSlide = (slide: Slide) => {
        const index = selectedSlideshow.slides.findIndex(s => s.position === slide.position);
        selectedSlideshow.slides[index] = slide;
        cli.updateSlideshow(selectedSlideshow).then((sl) => {
            setSelectedSlideshow(sl);
            setEditingSlide(undefined);
        });
    };
    const onEditClose = () => {
        setEditingSlide(undefined);
    };

    useEffect(() => {
        const st = location.state as { selectedSlideshow: Slideshow };
        setSelectedSlideshow(st.selectedSlideshow);
    }, [location.state]);

    return (
        <>
            <div>
                {selectedSlideshow ? <h2>{selectedSlideshow.name}</h2> : <h2>No slideshow selected...</h2>}
                {selectedSlideshow && selectedSlideshow.slides.length == 0 ? <h4>Slideshow is empty</h4> :
                    <Row xs={1} md={2} className={'g-4'}>
                        {
                            selectedSlideshow.slides.map((s) => {
                                return <SlideThumbnail slide={s} key={s.position} onRemove={handleRemoveSlide}
                                    onEdit={setEditingSlide} />;
                            })
                        }
                    </Row>
                }
                <Button variant="primary" onClick={handleAddSlide} style={{float: 'right'}}>Add slide</Button>
            </div>
            {editingSlide ?
                <EditSlideModal visible={true} onClose={onEditClose} onSave={handleSaveSlide} slide={editingSlide}
                    okButtonText={'Save'} /> : null}
        </>
    );
};
