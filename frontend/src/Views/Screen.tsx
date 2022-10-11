import React, {useContext, useEffect} from 'react';
import {ApiContext} from '../App';
import {useSearchParams} from 'react-router-dom';
import {ScreenPlayer} from '../Components/ScreenPlayer';

function Screen() {
    const cli = useContext(ApiContext);
    const [searchParams] = useSearchParams();
    const name = searchParams.get('name') ?? '';

    useEffect(() => {
        if (name !== '') {
            cli.registerScreen(name);
        }
    }, []);
    return (
        <div style={{
            backgroundColor: '#000000',
            overflow: 'hidden',
            cursor: 'auto',
            inset: '0px',
            position: 'absolute',
            width: '100%',
            height: '100%'
        }}>
            <ScreenPlayer name={name} />
        </div>
    );
}

export default Screen;