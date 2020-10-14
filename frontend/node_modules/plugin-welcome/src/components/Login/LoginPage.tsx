import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';

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

const initialUserState = {
  dOCTORNAME: 'นายแพทย์กำลัง นอนพอดี',
  dOCTOREMAIL: 'Gumlung@gmail.com',
};


export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบผู้เสียชีวิต' };
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [username, setUsername] = useState(String);

  const [loading, setLoading] = useState(true);

  const [useremail, setUserEmail] = useState(String);
  const [userid, setUserID] = useState(Number);
  const [loginstate, setLoginState] = useState(false);

  useEffect(() => {
    const getUsers = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
    };
    getUsers();
  }, [loading]);


  const UserEmailthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserEmail(event.target.value as string);
  };

  const LoginEmailChecker = async () => {
    setStatus(true);
    users.map((item:EntUser) =>{
      if (item.useremail === useremail){
        setUserID(item.id as number);
        setLoginState(true);
        setUsername(item.userName as string);
      }
      else {
        setLoginState(false);
      }
    }
    )
    if (loginstate == true) {
      setAlert(false);
    } else {
      setAlert(true);
    }
    const timer = setTimeout(() => {
      setStatus(false);
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
          <div>
            <Button variant="contained" color="primary">
              ออกจากระบบ
         </Button>
          </div>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  This is a success alert — check it out!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
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
              <TextField
                id="name"
                label="อีเมล์"
                variant="outlined"
                type="string"
                size="medium"
                value={useremail}
                onChange={UserEmailthandleChange}
                style={{ width: 400 }}
              />
            </FormControl>


            <div className={classes.margin}>
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
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}
