// Chakra imports
import { Avatar, Box, Flex, Text, useColorModeValue } from '@chakra-ui/react';
import Card from '../../features/Card';

export default function Banner(props: {
	banner: string;
    avatar: string;
    name: string;
    attendance_days: number | string;
    stay_time: number | string;
    number_of_coin: number | string;
    email: string;
    grade: string;
    [x: string]: any;
}) {
	const { banner, avatar, name, attendance_days, stay_time, number_of_coin, email, grade, ...rest } = props;
	// Chakra Color Mode
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
            <Text color={textColorSecondary} fontSize='md' fontWeight='400' mt='10px'>
                Email address: {email}
            </Text>
            <Text color={textColorSecondary} fontSize='md' fontWeight='400' mt='5px'>
                Rank: {grade}
            </Text>
			<Flex w='max-content' mx='auto' mt={{ base: 20, md: 150 }} >
				<Flex mx='auto' me={{ base: 5, md: 20 }} alignItems='center' flexDirection='column'>
					<Text color={textColorPrimary} fontSize='2xl' fontWeight='700'>
						{attendance_days}
					</Text>
					<Text color={textColorSecondary} fontSize='sm' fontWeight='400'>
						今月の出席日数
					</Text>
				</Flex>
				<Flex mx='auto' me={{ base: 5, md: 20 }} alignItems='center' flexDirection='column'>
					<Text color={textColorPrimary} fontSize='2xl' fontWeight='700'>
						{stay_time}
					</Text>
					<Text color={textColorSecondary} fontSize='sm' fontWeight='400'>
						今月の滞在時間
					</Text>
				</Flex>
				<Flex mx='auto' alignItems='center' flexDirection='column'>
					<Text color={textColorPrimary} fontSize='2xl' fontWeight='700'>
						{number_of_coin}
					</Text>
					<Text color={textColorSecondary} fontSize='sm' fontWeight='400'>
						保有コイン
					</Text>
				</Flex>
			</Flex>
		</Card>
	);
}