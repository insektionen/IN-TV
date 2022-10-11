import {Slide} from '../../API/Types';
import {ClockViewer} from './Clock/ClockViewer';

export const SlideViewer = (props: { slide: Slide, isNext: boolean, isCurrent: boolean }) => {
    switch (props.slide.type) {
    case 'picture':
    case 'text':
    case 'clock':
        return <ClockViewer slide={props.slide} />;
    default:
        return <h1>Unknown slide type</h1>;
    }
};