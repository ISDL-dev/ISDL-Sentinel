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
  Select,
} from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import { useUser } from '../../userContext';

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

const Grades = ['U4', 'M1', 'M2', 'D1', 'D2', 'D3', 'Teacher'] as const;
type Grade = typeof Grades[number];

interface SignUpData {
  name: string;
  auth_user_name: string;
  mail_address: string;
  password: string;
  grade_name: string;
}

export default function SignUp() {
  const [name, setName] = useState('');
  const [authUserName, setAuthUserName] = useState('');
  const [mailAddress, setMailAddress] = useState('');
  const [grade, setGrade] = useState<Grade>('U4');
  const [password, setPassword] = useState('');
  const toast = useToast();
  const navigate = useNavigate();
  const { setAuthUser } = useUser();

  const handleRegister = async () => {
    if (!name || !authUserName || !mailAddress || !password) {
      toast({
        title: 'Input required',
        description: 'Please fill in all fields',
        status: 'warning',
        duration: 5000,
        isClosable: true,
      });
      return;
    }

    const signUpData: SignUpData = {
      name,
      auth_user_name: authUserName,
      mail_address: mailAddress,
      password,
      grade_name: grade,
    };

    try {
      const response = await axios.post(`${baseURL}/sign-up`, signUpData);

      setAuthUser(response.data);
      toast({
        title: 'Registration successful',
        description: `Welcome, ${name}!`,
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
          errorMessage = 'User already exists.';
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
              Already have an account? <ChakraLink href="/sign-in-webauthn">Sign in</ChakraLink>
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
                <FormLabel htmlFor="name">Full Name</FormLabel>
                <Input
                  id="name"
                  type="text"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="authUserName">Username</FormLabel>
                <Input
                  id="authUserName"
                  type="text"
                  value={authUserName}
                  onChange={(e) => setAuthUserName(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="mailAddress">Email address</FormLabel>
                <Input
                  id="mailAddress"
                  type="email"
                  value={mailAddress}
                  onChange={(e) => setMailAddress(e.target.value)}
                />
              </FormControl>
              <FormControl>
                <FormLabel htmlFor="grade">Grade</FormLabel>
                <Select
                  id="grade"
                  value={grade}
                  onChange={(e) => setGrade(e.target.value as Grade)}
                >
                  {Grades.map((gradeOption) => (
                    <option key={gradeOption} value={gradeOption}>
                      {gradeOption}
                    </option>
                  ))}
                </Select>
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