<?php

namespace App\Models;

use App\Traits\DefaultDatetimeFormat;
use Illuminate\Contracts\Auth\MustVerifyEmail;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Illuminate\Support\Str;
use Laravel\Sanctum\HasApiTokens;
use Laravel\Sanctum\NewAccessToken;

/**
 * @property mixed $is_admin
 * @property mixed $activated_at
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

    /**
     * Create a new personal access token for the user.
     *
     * @param string $name
     * @param array $abilities
     * @param null $expiredAt
     * @return NewAccessToken
     */
    public function createToken(string $name, array $abilities = ['*'], $expiredAt = null): NewAccessToken
    {
        $token = $this->tokens()->create([
            'name' => $name,
            'token' => hash('sha256', $plainTextToken = Str::random(40)),
            'abilities' => $abilities,
            'expired_at' => now()->addMinutes($expiredAt),
        ]);

        return new NewAccessToken($token, $token->getKey().'|'.$plainTextToken);
    }
}
