import {Slide} from '../../../API/Types';
import {Thumbnail} from '../../Thumbnail';

export const ClockThumbnail = (props: { slide: Slide, onClick?: (s: Slide) => void, onRemoveClick?: () => void, onEditClick?: () => void }) => {
    return (
        <Thumbnail name={'Clock'} slide={props.slide} onClick={props.onClick} onRemoveClick={props.onRemoveClick}
            onEditClick={props.onEditClick}>
            <p>Clock</p>
        </Thumbnail>
    );
};