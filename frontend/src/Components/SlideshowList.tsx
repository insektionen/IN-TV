import React, {useContext, useEffect, useState} from 'react';
import {ApiContext} from '../App';
import {Dropdown} from 'react-bootstrap';
import {Loading} from './Loading';
import {Slideshow} from '../API/Types';
import {useNavigate} from 'react-router-dom';

export const SlideshowList = () => {
    const [slideshows, setSlideshows] = useState<Slideshow[]>([]);
    const cli = useContext(ApiContext);
    const navigate = useNavigate();

    const onSelect = (s: Slideshow) => {
        navigate('/edit', {state: {selectedSlideshow: s}});
    };

    useEffect(() => {
        cli.getSlideshows()
            .then(res => {
                setSlideshows(res);
            });
    }, []);

    return slideshows ?
        <>
            {slideshows.map((s) => {
                return <Dropdown.Item key={s.name} onClick={() => onSelect(s)}>{s.name}</Dropdown.Item>;
            })}
        </> : <Loading />;
};
