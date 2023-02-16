package repository

import "github.com/AlbatozK/go_backend_boilerplate/model"

type UserRepository struct {
	*GenericRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{GenericRepository: NewGenericRepository()}
}

func (gr *GenericRepository) GetUserRepository() *UserRepository {
	return &UserRepository{GenericRepository: gr}
}

func (ur *UserRepository) GetById(id int) (*model.User, error) {
	var err error
	if ur.tx == nil {
		err = ur.BeginTx()
		if err != nil {
			return nil, err
		}
		defer func() {
			if err != nil {
				ur.Rollback()
			} else {
				ur.Commit()
			}
		}()
	}
	user := &model.User{}
	err = ur.tx.Get(user, `SELECT * FROM "user" WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*model.User, error) {
	var err error
	if ur.tx == nil {
		err = ur.BeginTx()
		if err != nil {
			return nil, err
		}
		defer func() {
			if err != nil {
				ur.Rollback()
			} else {
				ur.Commit()
			}
		}()
	}
	user := &model.User{}
	err = ur.tx.Get(user, `SELECT * FROM "user" WHERE email=$1`, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) Create(user *model.User) error {
	var err error
	if ur.tx == nil {
		err = ur.BeginTx()
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				ur.Rollback()
			} else {
				ur.Commit()
			}
		}()
	}
	result, err := ur.tx.NamedExec(`
		INSERT INTO "user" (firstname, lastname, email) 
		VALUES (:firstname, :lastname, :email)
		RETURNING id
	`, user)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = int(userId)
	return nil
}

func (ur *UserRepository) Update(user *model.User) error {
	var err error
	if ur.tx == nil {
		err = ur.BeginTx()
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				ur.Rollback()
			} else {
				ur.Commit()
			}
		}()
	}
	_, err = ur.tx.NamedExec(`
		UPDATE "user" SET firstname=:firstname, lastname=:lastname, email=:email
		WHERE id=:id
	`, user)
	if err != nil {
		return err
	}
	return nil
}
