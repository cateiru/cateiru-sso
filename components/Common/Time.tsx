import {Text} from '@chakra-ui/react';
import React from 'react';
import {hawManyDaysAgo} from '../../utils/date';
import {Tooltip} from './Chakra/Tooltip';

interface Props {
  time: string | Date;
}

export const AgoTime: React.FC<Props> = ({time}) => {
  try {
    if (typeof time === 'string') {
      time = new Date(time);
    }

    return (
      <Tooltip placement="top" label={time.toLocaleString()}>
        <Text as="time" dateTime={time.toISOString()}>
          {hawManyDaysAgo(time)}
        </Text>
      </Tooltip>
    );
  } catch {
    return <Text>不正な時刻</Text>;
  }
};
