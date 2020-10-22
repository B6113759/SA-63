import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import Typography from '@material-ui/core/Typography';
import WelcomePage from '../WelcomePage';


import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);



export default function LoginPage() {
  const classes = useStyles();
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบผู้เสียชีวิต' };
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [LoginStatus, setLoginStatus] = useState(Boolean);
  const [alert, setAlert] = useState(Boolean);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [username, setUsername] = useState(String);

  const [loading, setLoading] = useState(true);

  const [useremail, setUserEmail] = useState(String);
  const [userid, setUserID] = useState(Number);

  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
    };
    getUsers();
  }, [loading]);


  const UserEmailthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserID(event.target.value as number);
  };

  const LoginEmailChecker = async () => {
    setStatus(true);
    const item = await api.getUser({ id: 1})
    setUsername(item.userName as string);
    setUserEmail(item.useremail as string);
    setAlert(true);
    setLoginStatus(true);

    const timer = setTimeout(() => {
      setStatus(false);
      setAlert(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      //subtitle="Some quick intro and links."
      ></Header>
      <Content>
        <ContentHeader title="ล็อคอินยืนยันตัวตน">
          <Typography align="left" style={{ marginRight: 16, color: "#00eeff" }}>
            {username}
          </Typography>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  เข้าสู่ระบบสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    อีเมล์ไม่ถูกต้อง
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel id="user-label">อีเมล์</InputLabel>
                <Select
                  labelId="user-label"
                  id="user"
                  value={userid}
                  onChange={UserEmailthandleChange}
                  style={{ width: 400 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.useremail}</MenuItem>
                  ))}
                </Select>
              </FormControl>


            <div className={classes.margin}>
            <Link to="/mainpage">
              <Button
                style={{ width: 400, backgroundColor: "#34eb77" }}
                color="primary"
                onClick={() => {
                  LoginEmailChecker();
                }}
                variant="contained"
              >
                ล็อคอิน
             </Button>
             </Link>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}
