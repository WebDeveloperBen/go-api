// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID           uuid.UUID  `json:"id"`
	AccountID    string     `json:"account_id"`
	ProviderID   string     `json:"provider_id"`
	UserID       uuid.UUID  `json:"user_id"`
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token"`
	IDToken      string     `json:"id_token"`
	ExpiresAt    *time.Time `json:"expires_at"`
	Password     string     `json:"password"`
}

type ActivityType struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Published bool      `json:"published"`
	Category  string    `json:"category"`
	CardColor string    `json:"card_color"`
}

type Asset struct {
	ID            uuid.UUID `json:"id"`
	FileName      string    `json:"file_name"`
	ContentType   string    `json:"content_type"`
	ETag          string    `json:"e_tag"`
	ContainerName string    `json:"container_name"`
	Uri           string    `json:"uri"`
	Size          int32     `json:"size"`
	Metadata      []byte    `json:"metadata"`
	IsPublic      bool      `json:"is_public"`
	Published     bool      `json:"published"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}

type Chapter struct {
	ID            uuid.UUID   `json:"id"`
	LessonID      uuid.UUID   `json:"lesson_id"`
	NavItemName   string      `json:"nav_item_name"`
	Description   string      `json:"description"`
	Order         pgtype.Int4 `json:"order"`
	Content       []byte      `json:"content"`
	Published     bool        `json:"published"`
	Title         string      `json:"title"`
	ChapterNumber pgtype.Int4 `json:"chapter_number"`
	UpdatedAt     time.Time   `json:"updated_at"`
	CreatedAt     time.Time   `json:"created_at"`
}

type Class struct {
	ClassID     uuid.UUID   `json:"class_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CourseCode  string      `json:"course_code"`
	Duration    pgtype.Int4 `json:"duration"`
	UpdatedAt   time.Time   `json:"updated_at"`
	CreatedAt   time.Time   `json:"created_at"`
}

type ClassScore struct {
	ClassID uuid.UUID `json:"class_id"`
	UserID  uuid.UUID `json:"user_id"`
	Score   int32     `json:"score"`
	Week    int32     `json:"week"`
}

type ClassUser struct {
	ClassID uuid.UUID  `json:"class_id"`
	UserID  *uuid.UUID `json:"user_id"`
}

type ClassesActivity struct {
	ID         uuid.UUID  `json:"id"`
	Term       string     `json:"term"`
	StartDate  *time.Time `json:"start_date"`
	EndDate    *time.Time `json:"end_date"`
	HomeworkID uuid.UUID  `json:"homework_id"`
	Day        string     `json:"day"`
	Period     string     `json:"period"`
	ClassID    uuid.UUID  `json:"class_id"`
	UpdatedAt  time.Time  `json:"updated_at"`
	CreatedAt  time.Time  `json:"created_at"`
}

type Course struct {
	CoursesID   uuid.UUID   `json:"courses_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Chapters    pgtype.Int4 `json:"chapters"`
	Sections    pgtype.Int4 `json:"sections"`
	Duration    pgtype.Int4 `json:"duration"`
	Category    string      `json:"category"`
	Image       string      `json:"image"`
	UpdatedAt   time.Time   `json:"updated_at"`
	CreatedAt   time.Time   `json:"created_at"`
}

type EducationalStandard struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Subject     string    `json:"subject"`
	Category    string    `json:"category"`
	SubCategory string    `json:"sub_category"`
	Level       int32     `json:"level"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type EducationsStandardsQuestionsMapping struct {
	QuestionID uuid.UUID `json:"question_id"`
	StandardID uuid.UUID `json:"standard_id"`
}

type Enrolment struct {
	OrganisationID *uuid.UUID `json:"organisation_id"`
	CourseID       *uuid.UUID `json:"course_id"`
	UserID         *uuid.UUID `json:"user_id"`
	Name           string     `json:"name"`
	Tags           []string   `json:"tags"`
	Content        []byte     `json:"content"`
	UpdatedAt      time.Time  `json:"updated_at"`
	CreatedAt      time.Time  `json:"created_at"`
}

type Group struct {
	GroupID   uuid.UUID  `json:"group_id"`
	Name      string     `json:"name"`
	CourseID  *uuid.UUID `json:"course_id"`
	Tags      []string   `json:"tags"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type GroupsUser struct {
	GroupID   *uuid.UUID `json:"group_id"`
	UserID    *uuid.UUID `json:"user_id"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type Homework struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	CreatedBy *uuid.UUID `json:"created_by"`
	Tags      []string   `json:"tags"`
	Content   []byte     `json:"content"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type HomeworkQuestion struct {
	ID         uuid.UUID  `json:"id"`
	QuestionID *uuid.UUID `json:"question_id"`
	HomeworkID *uuid.UUID `json:"homework_id"`
}

type HomeworkSubmission struct {
	ID         uuid.UUID   `json:"id"`
	UserID     *uuid.UUID  `json:"user_id"`
	HomeworkID *uuid.UUID  `json:"homework_id"`
	Progress   pgtype.Int4 `json:"progress"`
	Content    []byte      `json:"content"`
	UpdatedAt  time.Time   `json:"updated_at"`
	CreatedAt  time.Time   `json:"created_at"`
}

type Invitee struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	InvitedBy string    `json:"invited_by"`
	Relation  string    `json:"relation"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Jwk struct {
	ID         uuid.UUID        `json:"id"`
	PublicKey  string           `json:"public_key"`
	PrivateKey string           `json:"private_key"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type Lesson struct {
	LessonID       uuid.UUID   `json:"lesson_id"`
	SectionID      uuid.UUID   `json:"section_id"`
	CourseID       *uuid.UUID  `json:"course_id"`
	Title          string      `json:"title"`
	Summary        string      `json:"summary"`
	Content        []byte      `json:"content"`
	Duration       pgtype.Int4 `json:"duration"`
	OrderInSubject pgtype.Int4 `json:"order_in_subject"`
	Image          string      `json:"image"`
	UpdatedAt      time.Time   `json:"updated_at"`
	CreatedAt      time.Time   `json:"created_at"`
}

type LessonResource struct {
	LessonID   uuid.UUID `json:"lesson_id"`
	ResourceID uuid.UUID `json:"resource_id"`
}

type Module struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Asset       string    `json:"asset"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type Note struct {
	NoteID    uuid.UUID  `json:"note_id"`
	UserID    *uuid.UUID `json:"user_id"`
	Content   string     `json:"content"`
	LessonID  *uuid.UUID `json:"lesson_id"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type Organisation struct {
	OrganisationID uuid.UUID `json:"organisation_id"`
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	ContactDetails string    `json:"contact_details"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

type OrganisationUser struct {
	OrganisationID *uuid.UUID  `json:"organisation_id"`
	UserID         *uuid.UUID  `json:"user_id"`
	RoleID         pgtype.Int4 `json:"role_id"`
	UpdatedAt      time.Time   `json:"updated_at"`
	CreatedAt      time.Time   `json:"created_at"`
}

type Outcome struct {
	OutcomeID int32     `json:"outcome_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Presence struct {
	ID         uuid.UUID  `json:"id"`
	UserID     uuid.UUID  `json:"user_id"`
	LastStatus string     `json:"last_status"`
	LastLogin  *time.Time `json:"last_login"`
	LastLogout *time.Time `json:"last_logout"`
	UpdatedAt  time.Time  `json:"updated_at"`
	CreatedAt  time.Time  `json:"created_at"`
}

type Profile struct {
	ID            uuid.UUID   `json:"id"`
	Username      string      `json:"username"`
	Egl           pgtype.Int4 `json:"egl"`
	EnrolledAt    string      `json:"enrolled_at"`
	TermsAccepted bool        `json:"terms_accepted"`
	UpdatedAt     time.Time   `json:"updated_at"`
	CreatedAt     time.Time   `json:"created_at"`
}

type Question struct {
	ID                  uuid.UUID   `json:"id"`
	Egl                 pgtype.Int4 `json:"egl"`
	CurriculumReference []string    `json:"curriculum_reference"`
	CognitiveSkill      []string    `json:"cognitive_skill"`
	ReadingAbility      pgtype.Int4 `json:"reading_ability"`
	WritingAbility      pgtype.Int4 `json:"writing_ability"`
	ListeningAbility    pgtype.Int4 `json:"listening_ability"`
	Title               string      `json:"title"`
	Question            string      `json:"question"`
	Answer              string      `json:"answer"`
	Options             []byte      `json:"options"`
	Type                string      `json:"type"`
}

type Quiz struct {
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	ID             uuid.UUID   `json:"id"`
	TotalQuestions pgtype.Int4 `json:"total_questions"`
	UpdatedAt      time.Time   `json:"updated_at"`
	CreatedAt      time.Time   `json:"created_at"`
}

type QuizzesQuestion struct {
	ID            uuid.UUID   `json:"id"`
	QuizzesID     *uuid.UUID  `json:"quizzes_id"`
	Question      string      `json:"question"`
	Answer        string      `json:"answer"`
	Options       []byte      `json:"options"`
	Type          string      `json:"type"`
	QuestionOrder pgtype.Int4 `json:"question_order"`
	Resources     [][]byte    `json:"resources"`
	UpdatedAt     time.Time   `json:"updated_at"`
	CreatedAt     time.Time   `json:"created_at"`
}

type RandomQuestion struct {
	ID                  uuid.UUID   `json:"id"`
	Egl                 pgtype.Int4 `json:"egl"`
	CurriculumReference []string    `json:"curriculum_reference"`
	CognitiveSkill      []string    `json:"cognitive_skill"`
	ReadingAbility      pgtype.Int4 `json:"reading_ability"`
	WritingAbility      pgtype.Int4 `json:"writing_ability"`
	ListeningAbility    pgtype.Int4 `json:"listening_ability"`
	Title               string      `json:"title"`
	Question            string      `json:"question"`
	Answer              string      `json:"answer"`
	Options             []byte      `json:"options"`
	Type                string      `json:"type"`
}

type Resource struct {
	ResourceID   uuid.UUID `json:"resource_id"`
	ResourceType string    `json:"resource_type"`
	Size         int64     `json:"size"`
	Metadata     []byte    `json:"metadata"`
	Etag         string    `json:"etag"`
	Provider     string    `json:"provider"`
	Uri          string    `json:"uri"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Role struct {
	RoleID      int32     `json:"role_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type Section struct {
	SectionID   uuid.UUID   `json:"section_id"`
	CoursesID   uuid.UUID   `json:"courses_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Order       pgtype.Int4 `json:"order"`
	UpdatedAt   time.Time   `json:"updated_at"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Session struct {
	ID        uuid.UUID        `json:"id"`
	ExpiresAt pgtype.Timestamp `json:"expires_at"`
	IpAddress string           `json:"ip_address"`
	UserAgent string           `json:"user_agent"`
	UserID    uuid.UUID        `json:"user_id"`
}

type User struct {
	ID            uuid.UUID   `json:"id"`
	Fullname      string      `json:"fullname"`
	Email         string      `json:"email"`
	EmailVerified pgtype.Bool `json:"email_verified"`
	Image         string      `json:"image"`
	UpdatedAt     time.Time   `json:"updated_at"`
	CreatedAt     time.Time   `json:"created_at"`
}

type UserRole struct {
	ID          uuid.UUID  `json:"id"`
	UserID      *uuid.UUID `json:"user_id"`
	RoleID      int32      `json:"role_id"`
	ContextID   uuid.UUID  `json:"context_id"`
	ContextType string     `json:"context_type"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
}

type Verification struct {
	ID         uuid.UUID        `json:"id"`
	Identifier string           `json:"identifier"`
	Value      string           `json:"value"`
	ExpiresAt  pgtype.Timestamp `json:"expires_at"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

type WeeklyPlanner struct {
	ID        uuid.UUID `json:"id"`
	ClassID   uuid.UUID `json:"class_id"`
	Term      string    `json:"term"`
	Week      int32     `json:"week"`
	Content   []byte    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
