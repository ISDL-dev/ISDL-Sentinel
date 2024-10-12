import React, { useState } from 'react';
import axios from 'axios';
import {
  Box,
  Button,
  Container,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Stack,
  Text,
  useToast,
  Link as ChakraLink,
  RadioGroup,
  Radio,
  HStack,
} from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import { useUser } from '../../userContext';

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

type Grade = 'B4' | 'M1' | 'M2' | 'Teacher';

export default function SignUp() {
  const [fullName, setFullName] = useState('');
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [grade, setGrade] = useState<Grade>('B4');
  const [password, setPassword] = useState('');
  const toast = useToast();
  const navigate = useNavigate();
  const { setAuthUser } = useUser();

  const handleRegister = async () => {
    if (!fullName || !username || !email || !password) {
      toast({
        title: 'Input required',
        description: 'Please fill in all fields',
        status: 'warning',
        duration: 5000,
        isClosable: true,
      });
      return;
    }

    try {
      const response = await axios.post(`${baseURL}/register`, {
        fullName,
        username,
        email,
        grade,
        password,
      });

      setAuthUser(response.data);
      toast({
        title: 'Registration successful',
        description: `Welcome, ${fullName}!`,
        status: 'success',
        duration: 5000,
        isClosable: true,
      });
      navigate('/');

    } catch (error) {
      console.error('Registration error:', error);
      
      let errorMessage = 'An unexpected error occurred';
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 400) {
          errorMessage = 'Invalid input. Please check your information.';
        } else if (error.response?.status === 409) {
          errorMessage = 'Username or email already exists.';
        } else if (error.response?.status === 500) {
          errorMessage = 'Server error. Please try again later.';
        } else if (!error.response) {
          errorMessage = 'No response received from server';
        }
      }

      toast({
        title: 'Registration failed',
        description: errorMessage,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
    }
  };

  return (
    <Container maxW="lg" py={{ base: '32', md: '62' }} px={{ base: '0', sm: '8' }}>
      <Stack spacing="8">
        <Stack spacing="6">
          <Stack spacing={{ base: '2', md: '3' }} textAlign="center">
            <Heading size={{ base: '1xl', md: "2xl" }}>Create your account</Heading>
            <Text color="gray.600">
              Already have an account? <ChakraLink href="#">Log in</ChakraLink>
            </Text>
          </Stack>
        </Stack>
        <Box
          py={{ base: '8', sm: '8' }}
          px={{ base: '10', sm: '10' }}
          bg={{ base: 'white', sm: 'white' }}
          boxShadow={{ base: 'md', sm: 'md' }}
          borderRadius={{ base: 'xl', sm: 'xl' }}
        >
          <Stack spacing="6">
            <Stack spacing="5">
              <FormControl>
                <FormLabel htmlFor="fullName">Full Name</FormLabel>
                <Input
                  id="fullName"
                  type="text"
                  value={fullName}
                  onChange={(e) => setFullName(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="username">Username</FormLabel>
                <Input
                  id="username"
                  type="text"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="email">Email address</FormLabel>
                <Input
                  id="email"
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="grade">Grade</FormLabel>
                <RadioGroup onChange={(value) => setGrade(value as Grade)} value={grade}>
                  <HStack spacing="24px">
                    <Radio value="B4">B4</Radio>
                    <Radio value="M1">M1</Radio>
                    <Radio value="M2">M2</Radio>
                    <Radio value="Teacher">Teacher</Radio>
                  </HStack>
                </RadioGroup>
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="password">Password</FormLabel>
                <Input
                  id="password"
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </FormControl>
            </Stack>
            <Button colorScheme="teal" variant="solid" size="md" onClick={handleRegister}>
              Register
            </Button>
          </Stack>
        </Box>
      </Stack>
    </Container>
  );
}