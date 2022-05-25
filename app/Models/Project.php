<?php

namespace App\Models;

use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Database\Eloquent\Factories\HasFactory;

class Project extends Model
{
    use HasFactory, SoftDeletes;

    /**
     * The attributes that are mass assignable.
     *
     * @var string[]
     */
    protected $fillable = [
        'user_id', 'name',
    ];

    public function content()
    {
        return $this->morphOne(Content::class, 'contentable');
    }

    public function columns()
    {
        return $this->hasMany(ProjectColumn::class);
    }
}
