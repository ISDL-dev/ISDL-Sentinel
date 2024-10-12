import { Avatar, Box, Flex, Text, useColorModeValue } from '@chakra-ui/react';
import Card from '../../features/Card';

export default function Banner(props: {
  banner: string;
  avatar: string;
  name: string;
  attendance_days: number | string;
  stay_time: number | string;
  email: string;
  grade: string;
  roleList: string[] | null; // null の場合を考慮
  [x: string]: any;
}) {
  const { banner, avatar, name, attendance_days, stay_time, email, grade, roleList, ...rest } = props;

  const textColorPrimary = useColorModeValue('secondaryGray.900', 'white');
  const textColorSecondary = 'gray.400';
  const borderColor = useColorModeValue('white !important', '#111C44 !important');

  return (
    <Card mb={{ base: '0px', lg: '20px' }} alignItems='center' {...rest}>
      <Box bg={`${banner}`} bgSize='cover' borderRadius='16px' h='131px' w='100%' />
      <Avatar mx='auto' src={avatar} h='87px' w='87px' mt='-43px' border='4px solid' borderColor={borderColor} />
      <Text color={textColorPrimary} fontWeight='bold' fontSize='xl' mt='10px'>
        {name}
      </Text>

      {/* Centered Block with Aligned Labels */}
      <Box mt='10px' w='100%' maxW='400px' mx='auto'>
        {/* Email */}
        <Flex justifyContent='center' alignItems='baseline' flexWrap='wrap'>
          <Text color={textColorSecondary} fontSize='md' fontWeight='600' mr='2' textAlign='right' minW='120px'>
            Email address:
          </Text>
          <Text color={textColorPrimary} fontSize='md' fontWeight='400' wordBreak='break-word' flex='1'>
            {email}
          </Text>
        </Flex>

        {/* Grade */}
        <Flex justifyContent='center' alignItems='baseline' mt='5px' flexWrap='wrap'>
          <Text color={textColorSecondary} fontSize='md' fontWeight='600' mr='2' textAlign='right' minW='120px'>
            Grade:
          </Text>
          <Text color={textColorPrimary} fontSize='md' fontWeight='400' wordBreak='break-word' flex='1'>
            {grade}
          </Text>
        </Flex>

        {/* Roles */}
        {roleList && roleList.length > 0 && (
          <Flex justifyContent='center' alignItems='baseline' mt='5px' flexWrap='wrap'>
            <Text color={textColorSecondary} fontSize='md' fontWeight='600' mr='2' textAlign='right' minW='120px'>
              Roles:
            </Text>
            <Box flex='1'>
              {roleList.map((role, index) => (
                <Text key={index} color={textColorPrimary} fontSize='md' fontWeight='400'>
                  {role}担当
                </Text>
              ))}
            </Box>
          </Flex>
        )}
      </Box>

      {/* Attendance and Stay Time */}
      <Flex w='100%' justifyContent='center' mt={{ base: 20, md: 150 }}>
        <Flex mx='auto' me={{ base: 5, md: 20 }} alignItems='center' flexDirection='column'>
          <Text color={textColorPrimary} fontSize='2xl' fontWeight='700'>
            {attendance_days}
          </Text>
          <Text color={textColorSecondary} fontSize='sm' fontWeight='400'>
            今月の出席日数
          </Text>
        </Flex>
        <Flex mx='auto' alignItems='center' flexDirection='column'>
          <Text color={textColorPrimary} fontSize='2xl' fontWeight='700'>
            {stay_time}
          </Text>
          <Text color={textColorSecondary} fontSize='sm' fontWeight='400'>
            今月の滞在時間
          </Text>
        </Flex>
      </Flex>
    </Card>
  );
}
