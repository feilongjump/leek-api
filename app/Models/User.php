<?php

namespace App\Models;

use App\Traits\DefaultDatetimeFormat;
use Illuminate\Contracts\Auth\MustVerifyEmail;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Illuminate\Support\Facades\Mail;
use Laravel\Sanctum\HasApiTokens;
use App\Mail\ResetPassword;

/**
 * @property int $id
 * @property string $email
 * @property string $password
 * @property boolean $is_admin
 * @property \Carbon\Carbon $activated_at
 */
class User extends Authenticatable
{
    use HasApiTokens, HasFactory, Notifiable, DefaultDatetimeFormat;

    /**
     * The attributes that are mass assignable.
     *
     * @var array<int, string>
     */
    protected $fillable = [
        'name', 'email', 'password', 'activated_at'
    ];

    /**
     * The attributes that should be hidden for serialization.
     *
     * @var array<int, string>
     */
    protected $hidden = [
        'password',
        'remember_token',
    ];

    /**
     * The attributes that should be cast.
     *
     * @var array<string, string>
     */
    protected $casts = [
        'activated_at' => 'datetime',
    ];

    protected $appends = [
        'is_activated',
    ];

    public function getIsActivatedAttribute(): bool
    {
        return (bool) $this->activated_at;
    }

    public function sendPasswordResetNotification($token)
    {
        return Mail::to($this->email)->queue(new ResetPassword($this->email, $token));
    }
}
